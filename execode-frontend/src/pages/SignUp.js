import React, { useState } from 'react'
import UserPool from './UserPool';
// import { CognitoUserAttribute } from "amazon-cognito-identity-js";

function SignUp() {

    const [username, setUserName] = useState("");
    // const [phone, setPhone] = useState("+66");
    // const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    var attributeList = [];

    const onSubmit = (event) => {
        event.preventDefault();

        // var dataEmail = {
        //     Name: 'email',
        //     Value: email
        // }

        // var dataPhoneNumber = {
        //     Name: 'phone_number',
        //     Value: phone
        // }

        // var attributeEmail = CognitoUserAttribute(dataEmail);
        // var attributePhoneNumber = CognitoUserAttribute(dataPhoneNumber);

        // attributeList.push(attributeEmail);
        // attributeList.push(attributePhoneNumber);

        UserPool.signUp(username, password, attributeList, null, function(
            err,
            result
        ) {
            if (err) {
                alert(err.message || JSON.stringify(err));
                return;
            }
            var cognitoUser = result.user
            console.log('user name is ' + cognitoUser.getUsername());
        });
    }

    return (
        <div>
            <form onSubmit={onSubmit}>
                <label htmlFor='username'>Username</label>
                <input value={username} onChange={(event) => setUserName(event.target.value)}></input>
                {/* <label htmlFor='phonenumber'>Phone Number</label>
                <input value={phone} onChange={(event) => setPhone(event.target.value)}></input>
                <label htmlFor='email'>Email</label>
                <input value={email} onChange={(event) => setEmail(event.target.value)}></input> */}
                <label htmlFor='password'>Password</label>
                <input value={password} onChange={(event) => setPassword(event.target.value)}></input>
                <button type='submit'>Sign Up</button>
            </form>
        </div>
    )
}

export default SignUp