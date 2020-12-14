"use strict";
function error(msg, code) {
    throw { message: msg, code: code };
}
console.log(error('Something went wrong', 500));
