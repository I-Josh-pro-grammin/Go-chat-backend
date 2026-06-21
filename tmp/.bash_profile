# Since Application Control policies block running binaries from temp/cache folders,
# use this alias to compile and run your application directly in the workspace root:
alias gorun="go build -o main.exe main.go && ./main.exe"