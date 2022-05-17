import React from "react";
import NavbarComp from "../../components/Navbar/NavbarComp";
import { Form,InputGroup, FormControl, Button} from 'react-bootstrap';
import "./humidity.css";
import "bootstrap/dist/css/bootstrap.rtl.min.css";

const Humidity = () => {
    return (
        <>
            <NavbarComp />
            <span className="gap"></span>
            
            <div className="page-title">
                تنظیم رطوبت
            </div>
            <InputGroup className="mb-3 humidity-input" size="lg">
                <FormControl aria-label="First name" placeholder="نام دستگاه" />
            </InputGroup>
            <InputGroup dir="rtl" className="mb-3 humidity-input" size="lg">
                {/* <InputGroup.Text>تنظیم رطوبت</InputGroup.Text> */}
                <FormControl aria-label="min humidity" placeholder="حداقل رطوبت" />
                <FormControl aria-label="max humidity" placeholder="حداکثر رطوبت" />

                    <Button variant="outline-secondary" id="button-addon2" className="add-device-btn">
                    تنظیم
                    </Button>
            </InputGroup>
            <InputGroup dir="rtl" className="mb-3 humidity-input" size="lg">
            <Form type="date" className="date-form">
                <Form.Group controlId="formBasicEmail">
                    <Form.Label>تاریخ</Form.Label>
                    <Form.Control type="date" placeholder="تاریخ" />
                </Form.Group>
                <Form.Group controlId="formBasicPassword">
                    <Form.Label>ساعت</Form.Label>
                    <Form.Control type="time" placeholder="ساعت" />
                </Form.Group>
            </Form>
            <Form type="date" className="date-form">
                <Form.Group controlId="formBasicEmail">
                    <Form.Label>تاریخ</Form.Label>
                    <Form.Control type="date" placeholder="تاریخ" />
                </Form.Group>
                <Form.Group controlId="formBasicPassword">
                    <Form.Label>ساعت</Form.Label>
                    <Form.Control type="time" placeholder="ساعت" />
                </Form.Group>
            </Form>
            </InputGroup>
        </>
    );
}
export default Humidity;