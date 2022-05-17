import React, { useState } from 'react';
import Login from '../../pages/Login/Login';
import Home from '../../pages/Home'
import { Routes, Route, Redirect, Navigate } from "react-router-dom";
import useToken from './useToken';
import "../../index"
import Report from '../../pages/Report/Report';
import Humidity from '../../pages/Humidity/Humidity';
import Profile from '../../pages/Profile/Profile';
import Water from '../../pages/Water/Water';
import Devices from '../../pages/Devices/Devices';
import Warning from '../../pages/Warning/Warning';

function App() {
  const { token, setToken } = useToken();
  
  if(!token) {
    return (
    <Login setToken={setToken} />
    )
    
  }
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/login" element={<Login />} />
      <Route path="/profile" element={<Profile />} />
      <Route path="/report" element={<Report />} />
      <Route path="/humidity" element={<Humidity />} />
      <Route path="/water" element={<Water />} />
      <Route path="/devices" element={<Devices />} />
      <Route path="/warnings" element={<Warning />} />
      

    </Routes>
  );
}

export default App;
