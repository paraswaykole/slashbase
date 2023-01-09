import React, { FunctionComponent } from 'react'
import Projects from '../components/home/projects'
import WelcomeCard from '../components/home/welcomecard'
import { useAppSelector } from '../redux/hooks'

const HomePage: FunctionComponent<{}> = () => {

    return (
        <React.Fragment>
            {/* <WelcomeCard /> */}
            <Projects />
        </React.Fragment >
    )
}

export default HomePage
