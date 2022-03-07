const db = require('diskdb');

class LoginRepository {
    constructor() {
        this.GetUserByPhonePassword = this.GetUserByPhonePassword.bind(this);
    }

    GetUserByPhonePassword(phone, password){
        return db.users.find({phone: phone, password: password});
    }
}

module.exports = LoginRepository;