package lsp

import (
	"errors"
	"strings"

	"github.com/goccy/go-json"

	errors2 "arc/interpreter/errors"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"

	"github.com/tliron/commonlog"
	_ "github.com/tliron/commonlog/simple"

	"arc/interpreter"
)

const lsName = "my language"

var version string = "0.0.1"
var handler protocol.Handler

var log = commonlog.GetLogger("lsp")

type LanguageServer struct {
	*server.Server

	OpenDocuments map[string]*interpreter.SourceFile

	workspace    string
	parserErrors []errors2.ErrorPresenter
}

func Run(protocolMode string) {
	lsp := &LanguageServer{
		OpenDocuments: make(map[string]*interpreter.SourceFile),
	}
	lsp.boot(protocolMode)

}

func (self *LanguageServer) boot(protocolMode string) {
	// This increases logging verbosity (optional)
	p := "/Users/sam/Code/Projects/ArcLang/ArcInterpreter/lsp_logs.log"
	commonlog.Configure(1, &p)

	interpreter.Engine.DisableLogging(commonlog.GetWriter())

	handler = protocol.Handler{
		Initialize:          self.onInitialize,
		Initialized:         self.onInitialized,
		Shutdown:            self.onShutdown,
		SetTrace:            self.onSetTrace,
		TextDocumentDidOpen: self.onTextDocumentDidOpen,
		// TextDocumentDefinition: self.onTextDocumentDefinition,
		TextDocumentDidChange: self.onTextDocumentDidChange,
		TextDocumentDidClose:  self.onTextDocumentDidClose,
	}

	lspServer := server.NewServer(&handler, lsName, false)

	self.Server = lspServer

	var err error
	switch protocolMode {
	case "stdio":
		log.Infof("running stdio")
		err = lspServer.RunStdio()
	case "ws":
		log.Infof("running ws")
		err = lspServer.RunWebSocket("localhost:9090")
	case "tcp":
		log.Infof("running tcp")
		err = lspServer.RunTCP("localhost:9090")
	default:
		log.Criticalf("unknown protocol mode: %s", protocolMode)
	}

	if err != nil {
		log.Errorf("%s", err.Error())
	}
}

func (self *LanguageServer) onInitialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
	if params == nil {
		log.Infof("onInitialize - params: nil")
	} else {
		jsonData, _ := json.MarshalIndent(params, "", "  ")
		log.Infof("onInitialize - data: %s", string(jsonData))
	}

	interpreter.ErrorManager.AddProcessor(func(presenter *errors2.ErrorPresenter) error {
		self.Log.Errorf("Parser hook triggered: %s", presenter.Errors[0].Message)
		self.onParserError(presenter)
		return nil
	})

	if params != nil {
		if len(params.WorkspaceFolders) > 0 {
			self.workspace = params.WorkspaceFolders[0].URI
		} else {
			self.workspace = *params.RootPath
			log.Errorf("onInitialize - no workspace folder found")
		}
	}

	capabilities := handler.CreateServerCapabilities()
	capabilities.TextDocumentSync = protocol.TextDocumentSyncOptions{
		OpenClose: ptrVal(true),
		Change:    ptrVal(protocol.TextDocumentSyncKindFull),
	}

	initResult := protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    lsName,
			Version: &version,
		},
	}

	jsonData, _ := json.MarshalIndent(initResult, "", "  ")
	log.Infof("onInitialize(server initialize response) - data: %s", string(jsonData))

	return initResult, nil
}

func (self *LanguageServer) onInitialized(context *glsp.Context, params *protocol.InitializedParams) error {
	jsonData, _ := json.MarshalIndent(params, "", "  ")
	log.Infof("onInitialized - data: %s", string(jsonData))

	return nil
}

func (self *LanguageServer) onShutdown(context *glsp.Context) error {
	protocol.SetTraceValue(protocol.TraceValueOff)
	return nil
}

func (self *LanguageServer) onSetTrace(context *glsp.Context, params *protocol.SetTraceParams) error {
	protocol.SetTraceValue(params.Value)
	return nil
}

func (self *LanguageServer) onTextDocumentDidOpen(c *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	jsonData, _ := json.MarshalIndent(params, "", "  ")
	log.Infof("onTextDocumentDidOpen - data: %s", string(jsonData))

	script := &interpreter.SourceFile{
		Path:   params.TextDocument.URI,
		Source: params.TextDocument.Text,
	}

	self.OpenDocuments[params.TextDocument.URI] = interpreter.Engine.FinishLoadingScript(script)

	self.sendDiagnostics(
		c,
		params.TextDocument.URI,
	)

	interpreter.Engine.ProcessScripts()

	return nil
}

// func (self *LanguageServer) onTextDocumentDefinition(context *glsp.Context, params *protocol.DefinitionParams) (any, error) {
// 	jsonData, _ := json.MarshalIndent(params, "", "  ")
// 	log.Infof("onTextDocumentDefinition - data: %s", string(jsonData))
//
// 	doc, ok := self.OpenDocuments[params.TextDocument.URI]
// 	if !ok {
// 		log.Errorf("onTextDocumentDefinition - no document found for URI %s", params.TextDocument.URI)
// 		return nil, errors.New("no document loaded/parsed for URI: " + params.TextDocument.URI)
// 	}
//
// 	node := interpreter.TypeChecker.GetNodeAtPosition(doc.Program, params.Position)
// 	if node == nil {
// 		log.Errorf("onTextDocumentDefinition - no node found at position %d:%d", params.Position.Line, params.Position.Character)
// 		return nil, errors.New("no node found at position: " + params.TextDocument.URI)
// 	}
//
// 	nodeType, nodeTypeDeclaration := interpreter.TypeChecker.FindDeclaration(node)
// 	// nodeType := interpreter.TypeChecker.FindType(node)
// 	if nodeType == nil && nodeTypeDeclaration == nil {
// 		log.Errorf("onTextDocumentDefinition - no type found for node at position %d:%d", params.Position.Line, params.Position.Character)
// 		return nil, errors.New("no type found for node at position: " + params.TextDocument.URI)
// 	}
//
// 	nodeTypeDefinitionSource := interpreter.TypeChecker.GetNodeSourceFile(nodeType)
// 	if nodeTypeDefinitionSource == nil {
// 		log.Errorf("onTextDocumentDefinition - no source file found for type %s", nodeType.TypeName())
// 		return nil, errors.New("no source file found for type: " + params.TextDocument.URI)
// 	}
//
// 	response := []*protocol.Location{}
//
// 	nodeDeclTypePosRange := nodeType.GetTokenRange()
// 	if nodeDeclTypePosRange != nil {
// 		nodeDeclTypePosRange.ZeroIndexed()
// 		response = append(response, &protocol.Location{
// 			URI: nodeTypeDefinitionSource.Path,
// 			Range: protocol.Range{
// 				Start: protocol.Position{
// 					Line:      protocol.UInteger(nodeDeclTypePosRange.StartLine),
// 					Character: protocol.UInteger(nodeDeclTypePosRange.StartCol),
// 				},
// 				End: protocol.Position{
// 					Line:      protocol.UInteger(nodeDeclTypePosRange.StopLine),
// 					Character: protocol.UInteger(nodeDeclTypePosRange.RangeStopCol),
// 				},
// 			},
// 		})
// 	}
//
// 	nodeTypePosRange := nodeTypeDeclaration.GetTokenRange()
// 	if nodeTypePosRange != nil {
// 		nodeTypePosRange.ZeroIndexed()
//
// 		response = append(response, &protocol.Location{
// 			URI: nodeTypeDefinitionSource.Path,
// 			Range: protocol.Range{
// 				Start: protocol.Position{
// 					Line:      protocol.UInteger(nodeTypePosRange.StartLine),
// 					Character: protocol.UInteger(nodeTypePosRange.StartCol),
// 				},
// 				End: protocol.Position{
// 					Line:      protocol.UInteger(nodeTypePosRange.StopLine),
// 					Character: protocol.UInteger(nodeTypePosRange.RangeStopCol),
// 				},
// 			},
// 		})
// 	}
//
// 	resJsonData, _ := json.MarshalIndent(response, "", "  ")
// 	log.Infof("onTextDocumentDefinition - data: %s", string(resJsonData))
//
// 	return response, nil
// }

func (self *LanguageServer) onTextDocumentDidChange(context *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
	jsonData, _ := json.MarshalIndent(params, "", "  ")
	log.Infof("onTextDocumentDidChange - data: %s", string(jsonData))

	didApplySourceChanges := false
	if source, ok := self.OpenDocuments[params.TextDocument.URI]; ok {
		if len(params.ContentChanges) > 0 {
			change := params.ContentChanges[0]
			source.Source = change.(protocol.TextDocumentContentChangeEventWhole).Text
			didApplySourceChanges = true
		}
	} else {
		log.Errorf("onTextDocumentDidChange - no document found for URI %s", params.TextDocument.URI)
		return errors.New("no document loaded/parsed for URI: " + params.TextDocument.URI)
	}

	if didApplySourceChanges {
		interpreter.Engine.ProcessScripts()
		self.sendDiagnostics(
			context,
			params.TextDocument.URI,
		)

		self.Log.Infof("onTextDocumentDidChange - applied changes & re-processed scripts")
	} else {
		self.Log.Infof("onTextDocumentDidChange - no changes to apply")
	}

	return nil
}

func (self *LanguageServer) onTextDocumentDidClose(context *glsp.Context, params *protocol.DidCloseTextDocumentParams) error {
	jsonData, _ := json.MarshalIndent(params, "", "  ")
	log.Infof("onTextDocumentDidClose - data: %s", string(jsonData))
	return nil
}

func (self *LanguageServer) onParserError(presenter *errors2.ErrorPresenter) {
	self.parserErrors = append(self.parserErrors, *presenter)
}

func (self *LanguageServer) sendDiagnostics(c *glsp.Context, documentUri string) {
	if len(self.parserErrors) == 0 {
		return
	}

	diagnosticParams := &protocol.PublishDiagnosticsParams{
		URI:         documentUri,
		Diagnostics: make([]protocol.Diagnostic, 0),
	}

	for _, parserError := range self.parserErrors {
		for _, codeError := range parserError.Errors {

			diagnostic := protocol.Diagnostic{
				Range: protocol.Range{
					Start: protocol.Position{
						Line:      protocol.UInteger(codeError.Start.Line),
						Character: protocol.UInteger(codeError.HighlightBounds.StartColumn),
					},
					End: protocol.Position{
						Line:      protocol.UInteger(codeError.End.Line),
						Character: protocol.UInteger(codeError.HighlightBounds.EndColumn),
					},
				},
				Severity: ptrVal(protocol.DiagnosticSeverityError),
				Code: &protocol.IntegerOrString{
					Value: strToCamelCase(codeError.Message),
				},
				Source:  ptrVal("ArcLang"),
				Message: codeError.Message,
			}

			diagnosticParams.Diagnostics = append(diagnosticParams.Diagnostics, diagnostic)

		}
	}

	c.Notify("textDocument/publishDiagnostics", diagnosticParams)

	self.parserErrors = make([]errors2.ErrorPresenter, 0)
}

func ptrVal[T any](v T) *T {
	return &v
}

func strToCamelCase(str string) string {
	var result string

	words := strings.Split(str, "_")
	for _, word := range words {
		result += strings.Title(word)
	}

	return result
}
