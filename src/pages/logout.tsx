import { useAppDispatch } from '../redux/hooks'
import { logoutUser } from '../redux/currentUserSlice'
import { useEffect } from 'react'


const LogoutPage = () => {

    const dispatch = useAppDispatch()

   useEffect(()=>{
       const init = async () => {
        await dispatch(logoutUser())
       }
       init()
   }, [dispatch])

    return (
        <>

        <div className={"center-text"}>
            <h1>Logging out...</h1>
        </div>

        <style jsx>{`

                h1 {
                    font-size: 2.5rem;
                }

            `}</style>
        </>
    )
}

export default LogoutPage