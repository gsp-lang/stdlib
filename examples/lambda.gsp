(ns main
    "fmt"
    "github.com/jcla1/gisp/core"
    "github.com/eatonphil/gimpy/lambda")


(def main (fn []
    (let [[hd (lambda/car (lambda/cons 1 (2 nil)))]]
        (fmt/println hd))))
