import React from "react";
import {NotificationContainer, NotificationManager} from 'react-notifications';
import 'react-notifications/lib/notifications.css';
import axios from "axios";
import useToken from "../../components/App/useToken";
import NavbarComp from "../../components/Navbar/NavbarComp";

const Warning = () => {
    const url_getWarning = 'http://usagi.carriot.ir:8000/warning/get';
    const token = useToken();
    const createNotification = () => {
        NotificationManager.warning('Warning message', 'Close after 3000ms', 3000);
    };

    const getWarning = () => {
        axios.get(url_getWarning, {
            headers: {
                Authorization: 'Bearer ' + token
            }
        }).then(response => {

        })
    };

    return (
        <div>
            <NavbarComp />
        </div>
    );
}

export default Warning;