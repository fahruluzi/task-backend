module.exports = {
    checkBody: function (reqBody, requires, strings) {
        const required = "$_variable is required"
        const emailInvalid = "Email is invalid"

        let errors = {}
        for (let i in requires) {
            if (typeof reqBody[requires[i]] === 'undefined' || (reqBody[requires[i]] === "")) {
                errors[requires[i]] = required.replace('$_variable', requires[i])
            } else if (requires[i] === 'email' && !validateEmail(reqBody[requires[i]])) {
                errors[requires[i]] = emailInvalid
            }
        }

        if (Object.getOwnPropertyNames(errors).length > 0) {
            return errors
        }
        return false
    }
}

function validateEmail(email) {
    let re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(String(email).toLowerCase());
}
