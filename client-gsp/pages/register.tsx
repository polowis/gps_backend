import type { NextPage } from 'next'

const RegisterPage: NextPage = () => {
    return (
        <div>

        </div>
    )
}

export async function getServerSideProps() {
    const res: Response = await fetch(`${process.env.API_HOST}/auth/texture`, {
        method: "POST"
    })
    const data = await res.json()
  
    return { props: { data } }
}

export default RegisterPage;