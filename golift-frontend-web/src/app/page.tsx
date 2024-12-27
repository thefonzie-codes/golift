import { Button } from "@/components/ui/button"
import { Nav } from "@/components/nav"

export default function Home() {
  return (
    <>
      <Nav />
      <div className="container mx-auto px-4 py-8">
        <h1 className="text-4xl font-bold mb-8">Welcome to goLift 🏋🏻</h1>
        <div className="space-y-4">
          <Button>Get Started</Button>
        </div>
      </div>
    </>
  )
}
