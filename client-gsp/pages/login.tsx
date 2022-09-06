import type { NextPage } from 'next';
import dynamic from 'next/dynamic';
import React from 'react';
import Button from '../components/Button';
import TextureBox from '../components/non_canvas/TextureBox';
import styles from './register.module.css';

const TexturePlayground = dynamic(() => import("../components/canvas/TexturePlayground"), {
    ssr: false,
});

const BOX_LENGTH_MINIMUM = 1;

const LoginPage: NextPage = (data: any) => {
    const textureBoxRef = React.useRef();
    const [completeBox, setCompleteBox] = React.useState(false);
    const [savedOrder, setSavedOrder] = React.useState("");
    const [errorMessage, setErrorMessage] = React.useState("");
    const [session, setSession] = React.useState(data.data.session);
    const [email, setEmail] = React.useState("");

    const verify = () => {
        if (textureBoxRef.current) {
            const chosenOrders = textureBoxRef.current.getOrders()
            const orderArray = chosenOrders.split("_")
            if (orderArray.length < BOX_LENGTH_MINIMUM) { // must be at least 1 box chosen
                setErrorMessage("You haven't chosen your password yet!");
                return;
            }
            if (email.length < 6) { // garbage validation to be rework
                setErrorMessage("Email not set!");
                return;
            }
            setErrorMessage("") // clear error message
            //setCompleteBox(true);
            setSavedOrder(chosenOrders)
            sendVerifyRequest(chosenOrders)
        }
    }

    const sendVerifyRequest = (orders: any) => {
        const APIRequest = {
            "email": email,
            "order": orders
        }
        
        fetch(`${process.env.NEXT_PUBLIC_API_HOST}/auth/verify`, {
            method: "POST",
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(APIRequest)
        }).then(res => res.json()).then(result => {
            if (result.success) {
                window.location.href = result.redirect_url
            }
        })
    }
    return (
        <div className={styles['container']}>
            <h1 className={styles['h1']}>Login Demo</h1>
            <br/>
            <h1 className={styles['error']}>{errorMessage}</h1>
            {!completeBox ? 
                <input className={styles['input']} placeholder="Email" value={email} onChange={(e: React.ChangeEvent<HTMLInputElement>) => setEmail(e.target.value)}/>
            : null}
            <br/><br/>
            {completeBox ? 
                null
            : <React.Fragment>
                <TextureBox ref={textureBoxRef} data={data}/>
                <Button onClick={() => verify()}>Verify</Button>
              </React.Fragment>
            }
        </div>
           
    )
}

export default LoginPage;

export async function getServerSideProps() {
    const res: Response = await fetch(`${process.env.API_HOST}/auth/texture`, {
        method: "POST",
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
    })
    const data = await res.json()
  
    return { props: { data: data.data } }
}