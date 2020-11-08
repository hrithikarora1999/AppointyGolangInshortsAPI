package main
import "testing"
import "net/http/httptest"
import "net/http"



func TestGetEntries(t *testing.T) {
    req, err := http.NewRequest("GET", "/article/all", nil)
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(getBooks)
    handler.ServeHTTP(rr, req)
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected := `[{"_id":"5fa80ff3dc4d7c09eb422801","id":"1111","title":"Hrithik","Stitle":"HrithikArora","content":"I am the best","timestamp":"0001-01-01T00:00:00Z"},{"_id":"5fa817172640c3b7f2d5e9db","id":"1112","title":"Hrithik","Stitle":"HrithikArora","content":"I am the best","timestamp":"0001-01-01T00:00:00Z"},{"_id":"5fa817272640c3b7f2d5e9dc","id":"1113","title":"Hrithik","Stitle":"HrithikArora","content":"I am the best","timestamp":"0001-01-01T00:00:00Z"},{"_id":"5fa817922640c3b7f2d5e9de","id":"1114","title":"Dhruv","Stitle":"HrithikArora","content":"I am the best","timestamp":"0001-01-01T00:00:00Z"},{"_id":"5fa817a72640c3b7f2d5e9df","id":"1115","title":"Salman","Stitle":"Salman Khan","content":"I am the best","timestamp":"0001-01-01T00:00:00Z"},{"_id":"5fa81806369f95f0820cee82","id":"1116","title":"Tanya","Stitle":"Salman Khan","content":"I am the best","timestamp":"0001-01-01T00:00:00Z"},{"_id":"5fa81c4a8fec76eb13289a42","id":"1117","title":"Hrithik","Stitle":"HrithikArora","content":"I am the best","timestamp":"0001-01-01T00:00:00Z"},{"_id":"5fa81c822b0d3df94000258a","id":"1118","title":"Hrithik","Stitle":"HrithikArora","content":"I am the best","timestamp":"0001-01-01T00:00:00Z"},{"_id":"5fa8204123a556ea27e15f2c","id":"1119","title":"Hrithik","Stitle":"HrithikArora","content":"I am the best","timestamp":"0001-01-01T00:00:00Z"},{"_id":"5fa8316f1c98ec21bf7db117","id":"1120","title":"Hrithik","Stitle":"DhruvArora","content":"I am the best","timestamp":"0001-01-01T00:00:00Z"}]`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}


func TestGetSearchEntries(t *testing.T) {
    req, err := http.NewRequest("GET", "/article/test/search/q=Hrithik", nil)
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(SearchArticle)
    handler.ServeHTTP(rr, req)
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected := `[
        {
            "_id": "5fa8316f1c98ec21bf7db117",
            "id": "1120",
            "title": "Hrithik",
            "Stitle": "DhruvArora",
            "content": "I am the best",
            "timestamp": "0001-01-01T00:00:00Z"
        }
    ]`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}





