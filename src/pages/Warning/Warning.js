import React , {useState,useEffect} from "react";
import { ListGroup } from 'react-bootstrap';
import {NotificationContainer, NotificationManager} from 'react-notifications';
import 'react-notifications/lib/notifications.css';
import axios from "axios";
import useToken from "../../components/App/useToken";
import NavbarComp from "../../components/Navbar/NavbarComp";
import Popup from 'react-popup';
import "./warning.css";


const Warning = () => {
    const [warning, setWarning] = useState([]);
    const url_getWarning = "http://usagi.carriot.ir:8000/warning/get";
    const token = useToken().token;

    const getWarning = () => {
        console.log(token);
        axios.get(url_getWarning, {
            headers: {
                Authorization: 'Bearer ' + token
            }
        }).then(response => {
            console.log(response.data.warning)
            setWarning(response.data.warning);
        })
    };

    useEffect(() => {
        getWarning()
    }, []);

    return (
        <div>
            <NavbarComp />
            <span className="gap"></span>
            <div className="page-title">
                هشدارها
            </div>


            <ListGroup as="ol" numbered>
                {warning ? warning.map(item => {
                    return (
                        <>
                            <ListGroup.Item
                                as="li"
                                className="d-flex justify-content-between align-items-start"
                            >
                                <div className="ms-2 me-auto">
                                    <div className="fw-bold">
                                        {"سریال:\t"+item.device_serial}
                                    </div>
                                    <div>
                                        {"حد غیرمجاز:\t" + item.difference}
                                    </div>
                                    <div>
                                        {"زمان:\t" + item.time}
                                    </div>
                                </div>

                            </ListGroup.Item>


                        </>
                    );
                }) : <div className="warning">
                    شما اخیرا هشداری ندارید.
                </div>
                }
            </ListGroup>
            <Popup />
        </div>
    );
}

export default Warning;