(ns prelude)

(def Null (fn [z] ()))

(def Cons (fn [x y]
    (fn [m] (m x y))))

(def Car (fn [z]
    (z (fn [p q] p))))

(def Cdr (fn [z]
    (z (fn [p q] q))))

(def Nullq (fn [z]
    (= (Car z) ())))

(def Len (fn [z]
    (if (= (Nullq z) true)
        0
        (+ 1 (Len (Cdr z))))))

(def Map (fn [f z]
    (if (= (Nullq z) true)
        Null
        (Cons (f (Car z)) (Cdr z)))))

(def Append (fn [a z]
    (if (= (Nullq a) true)
        z
        (Cons (Car a) (Append (Cdr a) z)))))
