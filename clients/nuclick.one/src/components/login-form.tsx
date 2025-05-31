import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { useState } from "react"
import { useUser } from "@/modules/auth/hooks/useUser"

export function LoginForm({
    className,
    ...props
}: React.ComponentProps<"div">) {
    const [isLogin, setIsLogin] = useState(true)
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const [username, setUsername] = useState("")
    const [isLoading, setIsLoading] = useState(false)
    const [error, setError] = useState("")

    const { login, register } = useUser()

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault()
        setIsLoading(true)
        setError("")

        try {
            let success = false
            
            if (isLogin) {
                success = await login({
                    email,
                    password,
                    username: null
                })
            } else {
                success = await register({
                    username,
                    email,
                    password
                })
            }

            if (success) {
                // Redirect or update UI as needed
                console.log("Authentication successful!")
                window.location.reload()
            } else {
                setError("Authentication failed. Please try again.")
            }
        } catch (error) {
            console.error("Authentication error:", error)
            setError("An error occurred during authentication.")
        } finally {
            setIsLoading(false)
        }
    }

    return (
        <div className={cn("flex flex-col gap-6", className)} {...props}>
            <Card className="">
                <CardContent>
                    <h1 className="mb-10 text-2xl flex items-center">
                        <span className="text-[#CF082E]">
                            <img
                                src="brand/logo/logo.jpg"
                                className="mr-2 w-7"
                            />
                        </span>{" "}
                        vafaill
                    </h1>
                    <form onSubmit={handleSubmit}>
                        <div className="flex flex-col gap-4">
                            {!isLogin && (
                                <div className="grid gap-3">
                                    <Label htmlFor="username">Username</Label>
                                    <Input
                                        id="username"
                                        type="text"
                                        placeholder="johndoe"
                                        value={username}
                                        onChange={(e) => setUsername(e.target.value)}
                                        required={!isLogin}
                                    />
                                </div>
                            )}
                            <div className="grid gap-3">
                                <Label htmlFor="email">Email</Label>
                                <Input
                                    id="email"
                                    type="email"
                                    placeholder="johndoe@example.com"
                                    value={email}
                                    onChange={(e) => setEmail(e.target.value)}
                                    required
                                />
                            </div>
                            <div className="grid gap-3">
                                <div className="flex items-center">
                                    <Label htmlFor="password">Password</Label>
                                </div>
                                <Input
                                    id="password"
                                    type="password"
                                    required
                                    placeholder="***************"
                                    value={password}
                                    onChange={(e) => setPassword(e.target.value)}
                                />
                                {isLogin && (
                                    <a
                                        href="#"
                                        className="ml-auto inline-block underline text-sm underline-offset-4 hover:underline"
                                    >
                                        forgot password?
                                    </a>
                                )}
                            </div>
                            {error && (
                                <div className="text-red-500 text-sm text-center">
                                    {error}
                                </div>
                            )}
                            <div className="flex flex-col gap-3">
                                <Button 
                                    type="submit" 
                                    className="w-full"
                                    disabled={isLoading}
                                >
                                    {isLoading ? "Please wait..." : (isLogin ? "Login" : "Register")}
                                </Button>
                                <Button variant="secondary" className="w-full" type="button">
                                    {isLogin ? "Login" : "Register"} with Google
                                </Button>
                            </div>
                        </div>
                        <div className="mt-4 text-center text-sm">
                            {isLogin ? "Don't have an account?" : "Already have an account?"}{" "}
                            <button
                                type="button"
                                onClick={() => {
                                    setIsLogin(!isLogin)
                                    setError("")
                                    setEmail("")
                                    setPassword("")
                                    setUsername("")
                                }}
                                className="underline underline-offset-4 hover:text-primary"
                            >
                                {isLogin ? "Sign up" : "Sign in"}
                            </button>
                        </div>
                    </form>
                </CardContent>
            </Card>
        </div>
    )
}