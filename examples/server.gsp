(ns main
    "github.com/eatonphil/nucore/fmt"
    "github.com/eatonphil/nucore/net/http")

(def hello (fn [w r]
    (fmt/Fprintf w "hello")
    ()))

(def main (fn []
    (http/HandleFunc "/" hello)
    (http/ListenAndServe ":9090" nil)))
