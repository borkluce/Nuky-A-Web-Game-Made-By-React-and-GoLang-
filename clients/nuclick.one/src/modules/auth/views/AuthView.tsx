import { LoginForm } from "@/components/login-form"

interface AuthViewProps {
    className?: string
}

const AuthView: React.FC<AuthViewProps> = ({ className }) => {
    return <LoginForm className="w-[400px] mx-auto "></LoginForm>
}

export default AuthView
