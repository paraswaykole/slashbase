import React, { FunctionComponent, useEffect } from 'react'
import { healthCheck, selectAPIVersion } from '../../redux/apiSlice'
import { useAppDispatch, useAppSelector } from '../../redux/hooks'
import logo from '../../assets/images/logo-icon.svg'

const AboutPage: FunctionComponent<{}> = () => {

    const dispatch = useAppDispatch()

    const version = useAppSelector(selectAPIVersion)

    useEffect(() => {
        dispatch(healthCheck())
    }, [dispatch])

    return (
        <React.Fragment>

            <h1>About Slashbase</h1>
            <br />
            <img src={logo} width={44} height={50} /><br />
            <h2>Version </h2>
            <p>{version}</p>
            <span className="tinytext">Licensed under the Apache License 2.0</span><br />
            <span className="tinytext">Copyright © 2021-2023 Slashbase.com.</span>
        </React.Fragment>
    )
}

export default AboutPage
