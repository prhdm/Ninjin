import React, { useState } from 'react';
import Login from '../../pages/Login/Login';
import Home from '../../pages/Home'
import { Routes, Route } from "react-router-dom";
import useToken from './useToken';
import "../../index"

function App() {
  const { token, setToken } = useToken();
  
  if(!token) {
    return <Login setToken={setToken} />
  }
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/login" element={<Login />} />
    </Routes>
  );
}

export default App;
