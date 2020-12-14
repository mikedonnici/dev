function error(msg: string, code: number) {
    throw {message: msg,  code: code}
}

console.log(error('Something went wrong', 500))
