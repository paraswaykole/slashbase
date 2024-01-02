import React, { FunctionComponent } from 'react'
import { WelcomeCard, WelcomeCardServer } from '../components/home/welcomecard'
import Constants from '../constants'

const HomePage: FunctionComponent<{}> = () => {

    return (
        <React.Fragment>
            {Constants.Build === 'desktop' &&
                <WelcomeCard />
            }
            {Constants.Build === 'server' &&
                <WelcomeCardServer />
            }
        </React.Fragment >
    )
}

export default HomePage
