package main
import (
	"net/http"
	"fmt"
	"time"
	"strings"
	"html/template"
 	"github.com/gorilla/sessions"
)
//var globalSessions *session.Manager
var login string
type Dane struct{
	Title string
	User string

}
//type Session interface {
//	Set(key, value interface{}) error //set session value
//	Get(key interface{}) interface{}  //get session value
//	Delete(key interface{}) error     //delete session value
//	SessionID() string                //back current sessionID
//}
//type Provider interface {
//	SessionInit(sid string) (Session, error)
//	SessionRead(sid string) (Session, error)
//	SessionDestroy(sid string) error
//	SessionGC(maxLifeTime int64)
//}
//
//type Manager struct {
//	cookieName  string     //private cookiename
//	lock        sync.Mutex // protects session
//	provider    Provider
//	maxlifetime int64
//}
//var provides = make(map[string]Provider)

// Register makes a session provider available by the provided name.
// If a Register is called twice with the same name or if the driver is nil,
// it panics.
//func Register(name string, provider Provider) {
//	if provider == nil {
//		panic("session: Register provider is nil")
//	}
//	if _, dup := provides[name]; dup {
//		panic("session: Register called twice for provider " + name)
//	}
//	provides[name] = provider
//}
//
//func (manager *Manager) sessionId() string {
//	b := make([]byte, 32)
//	if _, err := io.ReadFull(rand.Reader, b); err != nil {
//		return ""
//	}
//	return base64.URLEncoding.EncodeToString(b)
//}
//
//func NewManager(provideName, cookieName string, maxlifetime int64) (*Manager, error) {
//	provider, ok := provides[provideName]
//	if !ok {
//		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
//	}
//	return &Manager{provider: provider, cookieName: cookieName, maxlifetime: maxlifetime}, nil
//}
//func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
//	manager.lock.Lock()
//	defer manager.lock.Unlock()
//	cookie, err := r.Cookie(manager.cookieName)
//	if err != nil || cookie.Value == "" {
//		sid := manager.sessionId()
//		session, _ = manager.provider.SessionInit(sid)
//		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxlifetime)}
//		http.SetCookie(w, &cookie)
//	} else {
//		sid, _ := url.QueryUnescape(cookie.Value)
//		session, _ = manager.provider.SessionRead(sid)
//	}
//	return
//}
var store = sessions.NewCookieStore([]byte("ala-ma-kota"))



func root(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Witam na mojej stronie", request.URL.Path)
	if request.URL.Path != "/"{
		http.Redirect(response, request, "/default", 302)
	} else {
		http.ServeFile(response, request, "logowanie.html")
	}
}
func logowanie(response http.ResponseWriter, request *http.Request) {
	if (request.Method == "POST") {
		var correctPassword string
		request.ParseForm()
		fmt.Println("login", request.Form["login"])
		//fmt.Println("pass", request.Form["pass"])

		correctPassword = "12345"
		correctLogin := "aziron"

		login = request.Form["login"][0]
		haslo := request.Form["pass"][0]

		//session, _ := store.Get(request, "logowanie")
		//session.Values["username"] = login

//		globalSessions = NewManager("memory","gosessionid",3600)
		/*sess := globalSessions.SessionStart(response, request)
		request.ParseForm()
		if request.Method == "GET" {
			t, _ := template.ParseFiles("login.gtpl")
			response.Header().Set("Content-Type", "text/html")
			t.Execute(response, sess.Get("username"))
		} else {
			sess.Set("username", request.Form["username"])
			http.Redirect(response, request, "/", 302)
		}
		*/

		if strings.EqualFold(login, correctLogin) && haslo == correctPassword {
			fmt.Println("Zalogowano")
			http.Redirect(response, request, "/internal", 302)
		} else {
			fmt.Println("Hasło lub login są niepoprawne")
			time.Sleep(2)
			http.Redirect(response, request, "/invalid_login", 302)
		}

	}


}
func internalHandler(response http.ResponseWriter, request *http.Request) {

	data := Dane{Title:"Witaj", User:login}
	//http.ServeFile(response, request, "templates/welcome.html")
	t,_ := template.ParseGlob("templates/*.html")
	t.ExecuteTemplate(response, "internal", data)
}
func invalidLogin(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "invalid_login.html")

}

func headerHandler(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "templates/header.html")

}

func footerHandler(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "templates/footer.html")

}

func welcomeHandler(response http.ResponseWriter, request *http.Request) {

	data := Dane{Title:"Witaj"}
	//http.ServeFile(response, request, "templates/welcome.html")
	t,_ := template.ParseGlob("templates/*.html")
	t.ExecuteTemplate(response, "welcome", data)

}
func defaultHandler(response http.ResponseWriter, request *http.Request) {

	data := Dane{Title:"Witaj"}
	//http.ServeFile(response, request, "templates/welcome.html")
	t,_ := template.ParseGlob("templates/*.html")
	t.ExecuteTemplate(response, "default", data)

}
func logoutFormHandler(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "templates/logoutForm.html")

}


func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/welcome", welcomeHandler)
	http.HandleFunc("/logowanie", logowanie)
	http.HandleFunc("/invalid_login", invalidLogin)
	http.HandleFunc("/header", headerHandler)
	http.HandleFunc("/footer", footerHandler)
	http.HandleFunc("/default", defaultHandler)
	http.HandleFunc("/internal", internalHandler)
	http.HandleFunc("/logoutForm", logoutFormHandler)
	http.HandleFunc("/favicon.ico", func (response http.ResponseWriter, request *http.Request) {})
	err := http.ListenAndServe("localhost:5555", nil)
	if (err != nil) {
		println(err)
	}



}

func init() {

}

// Na zrobienie zadań masz czas do czwartku do 20:30. W razie pytań pisz maila,
// pamiętaj o stworzeniu brancha jeśli zamierzasz robić commit kodu, który się nie kompiluje