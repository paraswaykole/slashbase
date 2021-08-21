import type { NextPage } from 'next'
import { useState } from "react"
import apiService from "../network/apiService"

const LoginPage: NextPage = ()=> {

    const [userEmail, setUserEmail] = useState('')
    const [userPassword, setUserPassword] = useState('')


    const onLoginBtn = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        let response = await apiService.loginUser(userEmail, userPassword)
        console.log(response)
    }

    return (
    <main>
        <div className="card card-container">
            <div className="card-content">
                <div className="content">
                    <h1>Login to Slashbase</h1>
                    <form onSubmit={onLoginBtn}>
                        <div className="field">
                            <label className="label">Email</label>
                            <div className="control has-icons-left">
                                <input 
                                    className="input" 
                                    type="email" 
                                    placeholder="Enter Email" 
                                    value={userEmail} 
                                    onChange={(e: React.ChangeEvent<HTMLInputElement>)=>{setUserEmail(e.target.value)}}
                                />
                                <span className="icon is-small is-left">
                                <i className="fas fa-envelope"></i>
                                </span>
                            </div>
                        </div>
                        <div className="field">
                            <label className="label">Password</label>
                            <div className="control has-icons-left">
                                <input 
                                    className="input" 
                                    type="password" 
                                    placeholder="Enter Password" 
                                    value={userPassword}
                                    onChange={(e: React.ChangeEvent<HTMLInputElement>)=>{setUserPassword(e.target.value)}}/>
                                <span className="icon is-small is-left">
                                <i className="fas fa-key"></i>
                                </span>
                            </div>
                        </div>
                        <div className="control">
                            <button className="button is-primary">Login</button>
                        </div>
                </form>
                </div>
            </div>
        </div>
        <style jsx>{`
            .card-container{
                max-width: 500px;
                margin: 100px auto 0px auto;
            }
         `}</style>
    </main>
    )
}

export default LoginPage