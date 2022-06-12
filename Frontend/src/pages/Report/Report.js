import React, { useState, useEffect, PureComponent } from "react";
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend, ResponsiveContainer } from 'recharts';
import { Form,InputGroup, FormControl, Button} from 'react-bootstrap';

import NavbarComp from "../../components/Navbar/NavbarComp";
import "./report.css";
import "bootstrap/dist/css/bootstrap.rtl.min.css";
import axios from 'axios';
import useToken from "../../components/App/useToken";

const Report = () => {
    const url_getDeviceLog = 'http://usagi.carriot.ir:8000/device/get_log';
    const url_getWateringLog = "http://usagi.carriot.ir:8000/watering-log/get";
    const [deviceLogs, setDeviceLog] = useState([])
    const [wateringLogs, setWateringLog] = useState([])
    const [inputBeginDate, setBeginDate] = useState('')
    const [inputEndDate, setEndDate] = useState('')
    const [inputSerial, setSerial] = useState('')
    const [currentDeviceSerial, setCurrentDeviceSerial] = useState(JSON.stringify({
        "device_serial": "869170034680885",
        "begin_time": "2022-05-01T03:25:54+00:00",
        "end_time": "2022-05-01T23:26:54+00:00"
    }))
    const token = useToken().token;

    const fetchData = (deviceSerial) => {
        if(deviceSerial == undefined){
            deviceSerial = currentDeviceSerial
        }
        axios.post(url_getDeviceLog, deviceSerial , {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        })
            .then((response) => {
                setDeviceLog(response.data.device_log)
                console.log(response)
            });
    };
    const fetchDataW = (deviceSerial) => {
        if(deviceSerial == undefined){
            deviceSerial = currentDeviceSerial
        }
        axios.post(url_getWateringLog, deviceSerial , {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        })
            .then((response) => {
                setWateringLog(response.data.wateringLog)
                console.log(response)
            });
    };

    useEffect(() => {
        fetchData()
        fetchDataW()
    }, []);

    const setDeviceSerial = (inputSerial, beginTime, endTime) =>{
        setCurrentDeviceSerial(JSON.stringify({
            "device_serial": inputSerial,
            "begin_time": beginTime + "T00:00:00+00:00",
            "end_time": endTime + "T23:59:59+00:00",
        }))
        fetchData(JSON.stringify({
            "device_serial": inputSerial,
            "begin_time": beginTime + "T00:00:00+00:00",
            "end_time": endTime + "T23:59:59+00:00",
        }))
        fetchDataW(JSON.stringify({
            "device_serial": inputSerial,
            "begin_time": beginTime + "T00:00:00+00:00",
            "end_time": endTime + "T23:59:59+00:00",
        }))
    }
    return (
        <>
            <NavbarComp />
            <span className="gap"></span>

            <div className="page-title">
                 گزارش ها
            </div>
            <Form type="date" className="date-form">
                <Form.Group controlId="formBasicEmail">
                    <Form.Label> سریال دستگاه</Form.Label>
                    <FormControl type="text" placeholder="سریال"
                                 onChange={ (event) => setSerial(event.target.value) }
                    />
                    <Form.Label>تاریخ شروع</Form.Label>
                    <Form.Control type="date" placeholder="تاریخ شروع"
                                  onChange={ (event) => setBeginDate(event.target.value) }
                    />
                    <Form.Label>تاریخ پایان</Form.Label>
                    <Form.Control type="date" placeholder="تاریخ پایان"
                                  onChange={ (event) => setEndDate(event.target.value) }
                    />
                </Form.Group>
                    <Button onClick={() => setDeviceSerial(inputSerial, inputBeginDate, inputEndDate)} variant="outline-secondary" id="button-addon2" className="add-device-btn view-report">
                        مشاهده
                    </Button>
            </Form>
        <ResponsiveContainer className="chartclass" width={"100%"} height={500}>
            <LineChart
                width={500}
                height={300}
                data={deviceLogs}
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
            </LineChart>
        </ResponsiveContainer>
        <ResponsiveContainer className="chartclas" width={"100%"} height={500}>
            <LineChart
                width={500}
                height={300}
                data={wateringLogs}
                margin={{
                    top: 5,
                    right: 30,
                    left: 20,
                    bottom: 5,
                }}
            >
                <CartesianGrid strokeDasharray="3 3" />
                <XAxis dataKey="time" />
                <YAxis />
                <Tooltip />
                <Legend />
                <Line type="monotone" dataKey="water_amount" stroke="#82ca9d" activeDot={{ r: 8 }} />
            </LineChart>
        </ResponsiveContainer>
            </>
    );
}

export default Report;