import React from "react";

const Button = (props) => {
    return (
        <button className="ui button">{props.name}</button>
    );
}

export default Button;