import React from "react";
import 'bootstrap/dist/css/bootstrap.css';
import "bootstrap/dist/css/bootstrap.rtl.min.css";
import { Container, Navbar, Nav, NavDropdown } from 'react-bootstrap';
import './navbar.css';
import UserAvatar from "../UserAvatar/UserAvatar";

const NavbarComp = () =>{
    
    const handleLogout = () => {
        localStorage.clear();
        window.location.reload();
    }

    return (
        <Navbar collapseOnSelect expand="lg" bg="dark" variant="dark" dir="rtl">
          <Container>
          {/* <Navbar.Collapse dir="rtl" id="responsive-navbar-nav">
            <Nav className="me-auto">
                 <UserAvatar />
                <Nav.Link href="/profile" className="edit-profile">پروفایل</Nav.Link>
            </Nav>
            </Navbar.Collapse> */}

          <Navbar.Collapse dir="rtl" id="responsive-navbar-nav">
            <Nav className="me-auto main-nav">
              {/* <Nav.Link href="/">داشبورد</Nav.Link> */}
              <Nav.Link href="/report">گزارش</Nav.Link>
              <Nav.Link href="/devices">دستگاه‌ها</Nav.Link>
              <Nav.Link href="/warnings">هشدارها</Nav.Link>
              {/* <NavDropdown title="تنظیمات" id="collasible-nav-dropdown"> */}
                <Nav.Link href="/humidity">تنظیمات رطوبت</Nav.Link>
                <Nav.Link href="/water">فرم آبدهی</Nav.Link>
              {/* </NavDropdown> */}
            </Nav>
            <Nav>
            <button className="btn logout" onClick={handleLogout}>خروج</button>
            </Nav>
          </Navbar.Collapse>

          <Navbar.Brand href="/" className="brand">
            USAGI
          </Navbar.Brand>
          </Container>
        </Navbar>
    );

}

export default NavbarComp;