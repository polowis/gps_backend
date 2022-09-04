import type { NextPage } from 'next';
import TexturePlayground from '../components/canvas/TexturePlayground';
import styles from './register.module.css';

const RegisterPage: NextPage = (data) => {
    return (
        <div className={styles['container']}>
            <h1 className={styles['h1']}>Register Demo</h1>
            <br/>
            <input className={styles['input']} placeholder="Email"/>
            <br/>
            <TexturePlayground data={data}/>
        </div>
    )
}

export async function getServerSideProps() {
    const res: Response = await fetch(`${process.env.API_HOST}/auth/texture`, {
        method: "POST"
    })
    const data = await res.json()
  
    return { props: { data: data.data } }
}

export default RegisterPage;