interface AuthViewProps {
    className?: string
}

const AuthView: React.FC<AuthViewProps> = ({ className }) => {
    return <div className={` ${className}`}></div>
}

export default AuthView
