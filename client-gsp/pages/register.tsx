import type { NextPage } from 'next';
import dynamic from 'next/dynamic';
import React from 'react';
import Button from '../components/Button';
//import TexturePlayground from '../components/canvas/TexturePlayground';
import TextureBox from '../components/non_canvas/TextureBox';
import styles from './register.module.css';

const TexturePlayground = dynamic(() => import("../components/canvas/TexturePlayground"), {
    ssr: false,
});

const SAFE_BOX_LENGTH = 6;

const RegisterPage: NextPage = (data: any) => {
    const textureBoxRef = React.useRef();
    const [completeBox, setCompleteBox] = React.useState(false);
    const [savedOrder, setSavedOrder] = React.useState("");
    const [errorMessage, setErrorMessage] = React.useState("");
    const [session, setSession] = React.useState(data.data.session);
    const [email, setEmail] = React.useState("");

    const save = () => {
        if (textureBoxRef.current) {
            const chosenOrders = textureBoxRef.current.getOrders()
            const orderArray = chosenOrders.split("_")
            if (orderArray.length < SAFE_BOX_LENGTH) {
                setErrorMessage("You must choose at least 6 boxes");
                return;
            }
            if (email.length < 6) { // garbage validation to be rework
                setErrorMessage("Email not set!");
                return;
            }
            setErrorMessage("") // clear error message
            setCompleteBox(true);
            setSavedOrder(chosenOrders)
            
        }
    }

    const handleSaveCanvas = (lines: any) => {
        register(lines)
    }

    const register= (lines: any) => {
        const APIRequest = {
            "order": savedOrder,
            "lines": lines,
            "email": email,
            "session": session
        }
        fetch(`${process.env.NEXT_PUBLIC_API_HOST}/auth/register`, {
            method: "POST",
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(APIRequest)
        })
    }
    return (
        <div className={styles['container']}>
            <h1 className={styles['h1']}>Register Demo</h1>
            <br/>
            <h1 className={styles['error']}>{errorMessage}</h1>
            {!completeBox ? 
                <input className={styles['input']} placeholder="Email" value={email} onChange={(e: React.ChangeEvent<HTMLInputElement>) => setEmail(e.target.value)}/>
            : null}
            <br/><br/>
            {completeBox ? 
                <TexturePlayground onSave={(lines) => handleSaveCanvas(lines)}/>
            : <React.Fragment>
                <TextureBox ref={textureBoxRef} data={data}/>
                <Button onClick={() => save()}>Save</Button>
              </React.Fragment>
            }
            
           
        </div>
    )
}

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

export default RegisterPage;