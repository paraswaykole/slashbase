import type { NextPage } from 'next'
import Head from 'next/head'
import { useRouter } from 'next/router'
import { useState } from "react"
import Constants from '../constants'
import { loginUser } from '../redux/currentUserSlice'
import { useAppDispatch } from '../redux/hooks'

const LoginPage: NextPage = ()=> {

    const [userEmail, setUserEmail] = useState('')
    const [userPassword, setUserPassword] = useState('')
    const [loginError, setLoginError] = useState(null)

    const dispatch = useAppDispatch()
    const router = useRouter()


    const onLoginBtn = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        try {
            await dispatch(loginUser({email: userEmail, password: userPassword})).unwrap()
            router.replace(Constants.APP_PATHS.HOME.path)            
        } catch (e: any){
            setUserPassword('')
            setLoginError(e)
        }
    }

    return (
    <main>
        <Head>
            <title>Login to Slashbase</title>
        </Head>
        <div className="card card-container">
            <div className="card-content">
                <div className="content">
                    <img src="/logo-icon.svg" width={44} height={50} />
                    <h1 className="heading1">Login to Slashbase</h1>
                    <form onSubmit={onLoginBtn}>
                        <div className="field">
                            <label className="label">Email</label>
                            <div className="control has-icons-left">
                                <input 
                                    className={`input${loginError ? ' is-danger':''}`} 
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
                                    className={`input${loginError ? ' is-danger':''}`} 
                                    type="password" 
                                    placeholder="Enter Password" 
                                    value={userPassword}
                                    onChange={(e: React.ChangeEvent<HTMLInputElement>)=>{setUserPassword(e.target.value)}}/>
                                <span className="icon is-small is-left">
                                    <i className="fas fa-key"></i>
                                </span>
                            </div>
                            {loginError && <span className="help is-danger">{loginError}</span> }
                        </div>
                        <br/>
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
            .heading1 {
                margin-top: 10px !important;
                margin-bottom: 20px !important;
            }
         `}</style>
    </main>
    )
}

export default LoginPage