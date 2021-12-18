import theme from "extendTheme";
import React from "react";
import LogoProps from "./Logo"

const PlateIcon: React.FC<LogoProps> = (props) => {
    var w = props.width ? props.width: "20px";
    var h = props.height ? props.height: "20px";
    var fill = props.fill ? props.fill : theme.colors.button;
    return (
        <svg width={w} height={h} fill="none" xmlns="http://www.w3.org/2000/svg">
            <path fill={fill} d="M12.44 2.40026C11.24 3.16026 10.76 3.24026 7.55999 2.96026C5.63999 2.84026 3.43999 2.80026 2.67999 2.96026C0.879988 3.28026 0.439988 4.36026 1.67999 5.36026C2.51999 6.04026 3.03999 6.08026 6.87999 5.84026C10.72 5.60026 11.28 5.68026 12.48 6.40026C14.2 7.44026 15.96 7.40026 17.64 6.28026C19.32 5.16026 19.52 4.04026 18.24 2.84026C16.84 1.52026 14.2 1.32026 12.44 2.40026ZM16.96 2.84026C18.72 3.64026 18.72 5.16026 16.96 5.96026C15.44 6.64026 13.84 6.48026 12.68 5.60026C11.72 4.84026 11.28 4.80026 7.11999 5.00026C2.43999 5.24026 0.679988 4.92026 1.99999 3.96026C2.39999 3.72026 4.19999 3.64026 7.11999 3.80026C11.28 4.00026 11.72 3.96026 12.68 3.20026C13.84 2.32026 15.44 2.16026 16.96 2.84026Z"/>
            <path fill={fill} d="M0.800049 10.4002C0.800049 13.3602 2.68005 16.4802 5.28005 17.8402C6.76005 18.6002 13.24 18.6002 14.72 17.8402C17.32 16.4802 19.2 13.3602 19.2 10.4002V8.80016H10H0.800049V10.4002ZM18.4 10.8402C18.4 12.5202 17.44 14.4002 15.76 16.0402L14.36 17.4002H10H5.64005L4.16005 16.0002C2.52005 14.4002 1.60005 12.6002 1.60005 10.8002V9.60016H10H18.4V10.8402Z"/>
        </svg>
    );
}


export default PlateIcon
