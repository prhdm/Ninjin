import React, { useState } from "react";
import Button from "../../components/Button";
import usagiLogo from '../../components/icons/icons8-carrot-60.png'
import PropTypes from 'prop-types';
import "./login.css";
import { useNavigate } from "react-router-dom";

async function loginUser(credentials) {
    return fetch( 'http://79.175.177.48:8000/login' , {
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
    const navigate = useNavigate()

    const handleSubmit = async e => {
        e.preventDefault();
        try {const response = await loginUser({
          username,
          password
        });
        console.log(response)
        if (response){setToken(response);
        navigate("/home")}

      }

        catch (err) {
          console.log(err)
        }
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