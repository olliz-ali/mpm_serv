// script.js
const HOME_URL="localhost:8080"
const TRAINING_URL="localhost:8080/Training"
const LOGIN_URL="localhost:8080/login"
const SIGNUP_URL="localhost:8080/signup"


class PageManager {
    constructor(HomeURL) {
        this.HomeURL=HomeURL;
    }


}

const PManager = new PageManager(HOME_URL);

document.getElementById("ai-coach-logo").setAttribute("href",HOME_URL);
document.getElementById("ref2Training").setAttribute("href",TRAINING_URL);
document.getElementById("loginbtn").setAttribute("href",LOGIN_URL);
document.getElementById("signupbtn").setAttribute("href",SIGNUP_URL);
