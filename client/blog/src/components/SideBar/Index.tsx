import "./index.scss";

export const Index = () => {
    const sideBarContents = [
        {
            title: 'UnRead',
        },
        {
            title: 'thread',
        }
    ];


    return (
        <div className="sidebar">
            {
                sideBarContents.map((item, index) => {
                    return (
                        <div className="sidebar-item" key={index}>
                            <div className="sidebar-item-title">
                                {item.title}
                            </div>
                        </div>
                    )
                })
            }
        </div>
    );
}
