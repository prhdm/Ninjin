import React from "react";
import profile from "./img/usagi pofile.jpg"
import 'bootstrap/dist/css/bootstrap.css';
import AvatarEditor from 'react-avatar-editor'
import './navbar.css';

const Navbar = () =>{
    
    const handleLogout = () => {
        localStorage.clear();
        window.location.reload();
    }

    return (
        <>
            <nav className="navbar navbar-dark bg-dark">
                <AvatarEditor className="avatar"
                    image={profile}
                    width={76}
                    height={76}
                    border={0}
                    color={[255, 255, 255, 0.6]} // RGBA
                    scale={1.2}
                    rotate={0}
                />
                <button className="btn logout" onClick={handleLogout}>logout</button>
            </nav> 

            
        </>
    );

}

export default Navbar;