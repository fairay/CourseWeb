import theme from "extendTheme";
import React from "react";
import LogoProps from "./Logo"

const OkIcon: React.FC<LogoProps> = (props) => {
    const {width="20px", height="20px", fill= theme.colors['accent-2'], ...rest} = props
    return (
        <svg width={width} height={height} fill="none" xmlns="http://www.w3.org/2000/svg">
            <path fill={fill} d="M10 0C4.486 0 0 4.486 0 10C0 15.514 4.486 20 10 20C15.514 20 20 15.514 20 10C20 4.486 15.514 0 10 0ZM10 18.1818C5.48848 18.1818 1.81818 14.5115 1.81818 10C1.81818 5.48855 5.48848 1.81818 10 1.81818C14.5115 1.81818 18.1818 5.48855 18.1818 10C18.1818 14.5115 14.5115 18.1818 10 18.1818Z"/>
            <path fill={fill} d="M13.7498 6.46448L8.60712 11.6071L6.25015 9.25006C5.89518 8.89509 5.31955 8.89503 4.96452 9.25C4.60948 9.60503 4.60948 10.1806 4.96452 10.5356L7.96427 13.5355C8.13476 13.706 8.36597 13.8018 8.60706 13.8018C8.60712 13.8018 8.60706 13.8018 8.60712 13.8018C8.84822 13.8018 9.07943 13.706 9.24991 13.5356L15.0354 7.75018C15.3905 7.39515 15.3905 6.81958 15.0354 6.46455C14.6804 6.10951 14.1048 6.10945 13.7498 6.46448Z"/>
        </svg>
    );
}

export default OkIcon