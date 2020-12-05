for i in {10..30}
do
  cd ./day$i
  echo "package main\n\nfunc main() {\n}" > day${i}.go
  echo "package main\n\nfunc TestMain() {\n}" > day${i}_test.go
  echo "" > input.txt
  cd ..
done
