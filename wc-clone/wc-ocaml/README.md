compile tests executable 
   ocamlfind ocamlc -o wc_test -package ounit2 -linkpkg wc.ml wc_test.ml

   ocamlc -o wc wc.ml


ocamlfind ocamlc -o wc_test -package oUnit -linkpkg -g wc.ml wc_test.ml
