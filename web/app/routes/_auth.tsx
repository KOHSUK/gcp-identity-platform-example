import { Outlet } from "@remix-run/react";

export default function AuthRoot() {
    return (
        <div>
            <Outlet />
        </div>
    )
}