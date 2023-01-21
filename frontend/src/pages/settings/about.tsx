
import React, { FunctionComponent } from 'react'
import { selectAPIVersion } from '../../redux/apiSlice'
import { useAppSelector } from '../../redux/hooks'
import logo from '../../assets/images/logo-icon.svg'

const AboutPage: FunctionComponent<{}> = () => {

    const version = useAppSelector(selectAPIVersion)

    return (
        <React.Fragment>

            <h1>About Slashbase</h1>
            <br />
            <img src={logo} width={44} height={50} /><br />
            <h2>Version </h2>
            <p>{version}</p>
            <span className="tinytext">Copyright © 2022 Slashbase.com.</span><br />
        </React.Fragment>
    )
}

export default AboutPage
