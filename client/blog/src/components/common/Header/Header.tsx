import "./Header.scss";
import {Cart} from './Icons/Cart'
import {List} from './Icons/List'

export const Header = () => {
    return (
        <>
            <div className={'header'}>
                <div className={'header__bs-list'}>
                    <List/>
                </div>
                <div className={'header__title'}>ECサイト</div>
                <div className={'header__cart'}>
                    <Cart/>
                </div>
            </div>
        </>
    )
}