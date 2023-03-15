**How To**
----
Clone https://github.com/i1-ns/tokenomy.git or download zip.

* **Run**

  Navigate your directory to github.com/i1-ns/tokenomy/cmd/backend then type go run main.go

* **Build:**

  go build -o `<binary file name>` `<path>`

* **Test**

  Navigate your directory to github.com/i1-ns/tokenomy/test the type go test endpoints_test.go -v

**API**
----
Returns json dummy data for id & name.

* **URL**

  api/v0/tokenomy

* **Method:**

  `GET`

* **Query Params:**
    * **Key:** id <br />
      **Value:** int <br />
      **Example:** ?id=1,3

* **Success Response:**

    * **Code:** 200 <br />
      **Status:** OK <br />
      **Content:** `{"code":200,"data":[{"id":1,"name":"A"},{"id":2,"name":"B"},{"id":3,"name":"C"}]}`

    * **Code:** 200 <br />
      **Status:** OK <br />
      **Content:** `{"code":200,"data":[{"id":1,"name":"A"},{"id":3,"name":"C"}]}`

* **Error Response:**

    * **Code:** 400 <br />
      **Status:** BAD REQUEST <br />
      **Content:** `{"code":400,"message":"strconv.Atoi: parsing \"A\": invalid syntax"}`

    * **Code:** 404 <br />
      **Status:** NOT FOUND <br />
      **Content:** `{"code":404,"message":"resource with ID [4] doesn't exists"}`

* **Sample Call:**

  ```javascript
    $.ajax({
      url: "api/v0/tokenomy",
      dataType: "json",
      type : "GET",
      success : function(r) {
        console.log(r);
      }
    });
  ```

**What's Next**

* **TODO**

    * **Generate godoc**

