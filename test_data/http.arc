
import "http"

object http {}


object SomeData {
  message string
  number  int
}

http {


  route GET "/injections/param/:name" {
	from body as data SomeData
	from route as name string
	data.message += " - name is " + name + " - " + string::to(2)
	return {"msg": data.message}
  }

  route GET "/injections/body" {
	from body as data SomeData
	data.message += "pls sir"
	return {"msg": data.message}
  }
  route GET "/hello" {
	return text "Hello world"
  }

  route GET "/hello/dict" {
	return { "hello": "world" } status 201
  }
  route GET "/hello/404" {
	return text "Not found!!!!" status 404
  }
  route GET "/hello/500/dict" {
	return { "error": "Something went wrong!!!" } status 500
  }

  route GET "/injections/query" {
	from body as data SomeData
	from query as name string
	data.message += " - name is " + name
	return {"msg": data.message}
  }

  route POST "/data/json" {
	from body as data SomeData

	return text "Hello world"
	return user;

//    http::get("..", RequestOptions{
//	  headers: {
//		"Content-Type": "application/json"
//	  }
//	})

	return text "Not found" status 404
	return status 404
  }
}

func main() {}