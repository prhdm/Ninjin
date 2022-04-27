import React, { useState } from "react";
import Button from "../../components/Button";
import usagiLogo from '../../components/icons/icons8-carrot-60.png'
import PropTypes from 'prop-types';
import "./login.css";

async function loginUser(credentials) {
    return fetch('http://localhost:8000/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(credentials)
    })
      .then(data => data.json())
      .then(response => {

        if(response.status && response.token){  
        localStorage.setItem("jwt", response.token)}})
   }

const Login = ({ setToken }) => {

    const [username, setUserName] = useState();
    const [password, setPassword] = useState();

    const handleSubmit = async e => {
        e.preventDefault();
        const token = await loginUser({
          username,
          password
        });
        console.log("mamd",token)
        setToken(token);
    }

    return (
        <React.Fragment>            
            <form className="ui form" onSubmit={handleSubmit}>
                <h1>ninjin</h1>
                
                <img className="logo ninjin" src={usagiLogo} alt="Ninjin logo. logo"/>
                <div className="field">
                    <label>Username</label>
                    <input type="text" name="Username" placeholder="Username" onChange={e => setUserName(e.target.value)}/>                   
                </div>

                <div className="field">
                    <label>Password</label>
                    <input type="password" name="password" placeholder="Password" onChange={e => setPassword(e.target.value)}/>
                </div>

                <div className ="field">
                    
                </div>
                <Button type="submit"  name="Login"/>
                
                
            </form>

        </React.Fragment>);
}


export default Login;

Login.propTypes = {
    setToken: PropTypes.func.isRequired
  }