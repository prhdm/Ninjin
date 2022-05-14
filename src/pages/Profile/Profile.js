import React from "react";
import NavbarComp from "../../components/Navbar/NavbarComp";
import { Form, Button } from 'react-bootstrap';
import "bootstrap/dist/css/bootstrap.rtl.min.css";
import UserAvatar from "../../components/UserAvatar/UserAvatar";
import "./profile.css";
const Profile = () => {
    return (
        <>
            <NavbarComp />
            <Form dir="rtl">
                <div  className ="profile-edit">
                    <UserAvatar />
                </div>
                <span className="gap"></span>
            <Form.Group className="mb-3" controlId="formBasicPassword">
                <Form.Label>نام</Form.Label>
                <Form.Control type="name" placeholder="Name" />
            </Form.Group>

            <Form.Group className="mb-3" controlId="formBasicEmail">
                <Form.Label>ایمیل</Form.Label>
                <Form.Control dir="rtl" type="email" placeholder="Enter email" />
            </Form.Group>

            <span className="gap"></span>
            <div>
                <Button variant="primary" className="login edit btn" size="sm">
                    ویرایش
                </Button>{' '}
                <Button href="/" variant="secondary" className="btn" size="sm">
                    بازگشت
                </Button>
            </div>
            </Form>
        </>
    );
}

export default Profile;