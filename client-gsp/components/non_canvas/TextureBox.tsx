import Image from 'next/image';
import React from 'react';
import Button from '../Button';
import styles from './texture.module.css';



const TextureBox = React.forwardRef(({data}: any, ref)=> {
    const [order, setOrder] = React.useState("");

    React.useImperativeHandle(ref, () => ({

        getOrders(): String {
            return order;
        }
    
    }));
    const handleBoxSelect = (id: string, index: number) => {
        
        if (order.length === 0) { // inital add
            setOrder(index.toString())
        } else {
            let tempOrder = (' ' + order).slice(1);
            tempOrder += "_"
            tempOrder += index.toString()
            setOrder(tempOrder)
        }
    }

    const handleClearButton = () => {
        setOrder("")
    }
    return (
        <div>
            {order.length > 0 ? 
                <div className={styles['result']}>
                    You've selected: &nbsp;
                    {order.split("_").map((order, index) => 
                        <React.Fragment key={index}>
                            {order}
                        </React.Fragment>
                    )}
                </div>
            : null}
            
            <div className={styles['choose-title']}>
                Choose your pattern! (image does not matter)
            </div>
            <div>
                <Button onClick={() => handleClearButton()}>Clear</Button>
            </div>
            <div className={styles['b']}>
            {data.data.images.map((box: any, index: number) => 
                <div className={styles['container']} key={index} onClick={() => handleBoxSelect(box.id, index)}>
                    <Image src={box.url} width={40} height={40}/>
                </div>
            )}
            </div>
        </div>
        
    )
})

export default TextureBox;