import React from "react";
import 'bootstrap/dist/css/bootstrap.css';
import "bootstrap/dist/css/bootstrap.rtl.min.css";
import { Container, Navbar, Nav } from 'react-bootstrap';
import './navbar.css';

const NavbarComp = () =>{
    
    const handleLogout = () => {
        localStorage.clear();
        window.location.reload();
    }

    return (
        <Navbar collapseOnSelect expand="lg" bg="dark" variant="dark" dir="rtl">
          <Container>

          <Navbar.Collapse dir="rtl" id="responsive-navbar-nav">
            <Nav className="me-auto main-nav">
              <Nav.Link href="/">دستگاه‌ها</Nav.Link>
                <Nav.Link href="/report">گزارش</Nav.Link>
                <Nav.Link href="/warning">هشدارها</Nav.Link>
                <Nav.Link href="/humidity">تنظیمات رطوبت</Nav.Link>
                <Nav.Link href="/water">فرم آبدهی</Nav.Link>
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