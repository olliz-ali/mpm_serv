// script.js
const HOME_URL="localhost:8080"
const TRAINING_URL="localhost:8080/Training"
const LOGIN_URL="localhost:8080/login"
const SIGNUP_URL="localhost:8080/signup"


class SignUp {
    constructor(HomeURL) {
        this.HomeURL=HomeURL;
    }
    emailContinue(){
        this.email=document
    }

}

const PManager = new SignUp(HOME_URL);

//hrefs
document.getElementById("log").setAttribute("href",HOME_URL)
document.getElementById("to-login").setAttribute("href",LOGIN_URL)

//connect events

document.getElementById("continue-btn").addEventListener("click",() => {
    PManager.emailContinue();
})
