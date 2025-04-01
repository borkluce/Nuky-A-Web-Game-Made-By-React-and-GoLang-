import React from "react"

// Components
import CLogoSection from "../../core/components/CLogoSection"

// Icons
import { FaBrain } from "react-icons/fa"
import { RiGitRepositoryFill } from "react-icons/ri"
import { GiStairsGoal } from "react-icons/gi"
import { MdAlternateEmail } from "react-icons/md"

import { Link } from "react-router-dom"
import { IconType } from "react-icons"

interface NavbarProps {
    isHorizontal: boolean
}

const Navbar: React.FC<NavbarProps> = ({ isHorizontal }) => {
    const renderLinkButton = (
        endpointWithoutSlash: string,
        title: string,
        subtitle?: string,
        IconComponent?: IconType
    ) => {
        return (
            <Link
                to={endpointWithoutSlash}
                className={`
                            flex justify-start items-center hover:text-gray-400 py-4 px-4
                                ${
                                    document.URL.endsWith(
                                        "/" + endpointWithoutSlash
                                    )
                                        ? "bg-1 bg-opacity-10"
                                        : ""
                                }
                            `}
            >
                <div
                    className={`flex flex-col items-end justify-center ${
                        IconComponent && "mr-4"
                    }`}
                >
                    <p className="">{title}</p>
                    <p className="text-sm text-white/60">{subtitle}</p>
                </div>
                {IconComponent && <IconComponent className="text-2xl" />}
            </Link>
        )
    }

    return (
        <nav
            className={`
                  text-white z-40
                    sticky top-0
                ${isHorizontal ? "drop-shadow-sm bg-black" : ""}
            `}
        >
            <div
                className={`
                    drop-shadow-lg
                    mx-auto flex
                
                    ${
                        isHorizontal
                            ? "container justify-between items-center"
                            : "h-full flex-col pl-16 pr-4 text-nowrap py-6 z-10 items-end text-end "
                    }
                `}
            >
                <div className="px-4">
                    <CLogoSection isHorizontal={isHorizontal} lightMode />
                </div>

                {!isHorizontal && (
                    <hr className="w-full my-8 border-black opacity-20" />
                )}

                <div
                    className={`
                        text-end
                        items-end
                flex h-full
                ${isHorizontal ? "space-x-12" : "flex-col space-y-4"}
                `}
                >
                    {renderLinkButton(
                        "project",
                        "Projects",
                        "apps & services",
                        RiGitRepositoryFill
                    )}
                    {renderLinkButton(
                        "about-us",
                        "About Us",
                        "team behind",
                        FaBrain
                    )}
                    {renderLinkButton(
                        "objective",
                        "Objectives",
                        "our vision",
                        GiStairsGoal
                    )}
                    {renderLinkButton(
                        "contact",
                        "Contact",
                        "contact us",
                        MdAlternateEmail
                    )}
                </div>
            </div>
        </nav>
    )
}

export default Navbar
