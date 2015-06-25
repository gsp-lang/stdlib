(ns main
    "github.com/jcla1/gisp/core"
    "github.com/eatonphil/gimpy/fmt"
    "github.com/eatonphil/gimpy/net/http")

(def hello (fn [w r]
    (fmt/Fprintf w "hello")
    ()))

(def main (fn []
    (http/HandleFunc "/" hello)
    (http/ListenAndServe ":9090" nil)))
