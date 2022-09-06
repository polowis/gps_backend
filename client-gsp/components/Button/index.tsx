import React from 'react'
import styles from './button.module.css'

type Props = {
    children: React.ReactNode
    onClick: React.MouseEventHandler<HTMLDivElement>
};

export default function Button({ children, onClick }: Props ){
    return (
        <div className={styles['button']} onClick={onClick}>
            {children}
        </div>
    )
}