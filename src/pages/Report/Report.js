import React, { useState, useEffect, PureComponent } from "react";
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend, ResponsiveContainer } from 'recharts';

import NavbarComp from "../../components/Navbar/NavbarComp";
import { ListGroup, InputGroup, FormControl, Button} from 'react-bootstrap';
import "./report.css";
import "bootstrap/dist/css/bootstrap.rtl.min.css";
import axios from 'axios';
import useToken from "../../components/App/useToken";

const Report = () => {
    const url_getDeviceLog = 'http://usagi.carriot.ir:8000/device/get_log';
    //const url_getWateringLog = "http://usagi.carriot.ir:8000/device/delete_device";
    const [devicelogs, setdevicelog] = useState([])
    const [deviceSerial, setdeviceserial] = useState({
        device_serial: "869170034680885",
        begin_time: "2022-05-17T13:25:54+02:00",
        end_time: "2022-05-17T13:26:54+02:00",
    })
    const token = useToken().token;

    const fetchData = () => {
        axios.post(url_getDeviceLog, deviceSerial , {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        })
            .then((response) => {
                setdevicelog(response.data.device_log)
                console.log(response)
            });
    };

    useEffect(() => {
        fetchData()
    }, []);

    const setDeviceSerial = (deviceSerial, beginTime, endTime) =>{
        setdeviceserial(deviceSerial, beginTime, endTime)
        fetchData()
    }
    return (
        <>
            <NavbarComp />
            <span className="gap"></span>

            <div className="page-title">
                 گزارش ها
            </div>
        <ResponsiveContainer width={"90%"} height={500}>
            <LineChart
                width={500}
                height={300}
                data={devicelogs}
                margin={{
                    top: 5,
                    right: 30,
                    left: 20,
                    bottom: 5,
                }}
            >
                <CartesianGrid strokeDasharray="3 3" />
                <XAxis dataKey="datetime" />
                <YAxis />
                <Tooltip />
                <Legend />
                <Line type="monotone" dataKey="humidity" stroke="#8884d8" activeDot={{ r: 8 }} />
                {/*<Line type="monotone" dataKey="uv" stroke="#82ca9d" />*/}
            </LineChart>
        </ResponsiveContainer>
            </>
    );
}

export default Report;