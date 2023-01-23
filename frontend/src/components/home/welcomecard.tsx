import React, { FunctionComponent } from 'react'
import { useAppDispatch } from '../../redux/hooks'
import logo from '../../assets/images/logo-icon.svg'



const WelcomeCard: FunctionComponent<{}> = () => {

    const dispatch = useAppDispatch()

    return (
        <React.Fragment>
            <div className='card'>
                <div className="card-content">
                    <img src={logo} width={45} alt="slashbase logo" />
                    <h1>Welcome to Slashbase!</h1>
                    <hr />
                    <div>
                        <h3>Loving or hating Slashbase?</h3>
                        <br />
                        <a href="https://discord.gg/U6fXgm3FAX" target="_blank" rel="noreferrer">
                            <button className="button">
                                <i className={`fab fa-discord`} />&nbsp; Share your feedback on our discord server.
                            </button>
                        </a>
                    </div>
                </div>
            </div>
        </React.Fragment >
    )
}


export default WelcomeCard