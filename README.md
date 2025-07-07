
![Penny Black](.readme/logo.png)

# Penny Black

This is a project to build an e-ink reader for the Raspberry Pi zero.
It's a hobby project to learn about Go, e-ink displays and Raspberry Pi.
also I want to use a dial to navigate the UI.

```mermaid
graph TB
    subgraph "Development Environment"
        Dev[Developer] --> FyneApp[Fyne UI App]
    end
    
    subgraph "Virtual Display System"
        Xvfb[Xvfb Virtual X Server<br/>:99 Display]
        FyneApp --> Xvfb
    end
    
    subgraph "UI Generation Process"
        FyneApp --> Lists[Lists & Routes UI]
        Lists --> Screenshot[Screenshot Timer<br/>Every 2 seconds]
        Screenshot --> ImageMagick[ImageMagick<br/>import command]
        ImageMagick --> Xvfb
    end
    
    subgraph "File System"
        ImageMagick --> Folder[screendump/<br/>current.png]
        Folder --> FileWatch[File Watcher<br/>fsnotify]
    end
    
    subgraph "E-ink Display System"
        FileWatch --> EinkApp[E-ink Display App]
        EinkApp --> ImageProcess[Image Processing<br/>• Format conversion<br/>• Dithering<br/>• Resize]
        ImageProcess --> EinkDriver[E-ink Hardware Driver]
        EinkDriver --> EinkScreen[E-ink Display]
    end
    
    subgraph "User Interaction"
        EinkScreen --> User[User Input]
        User --> InputHandler[Input Handler]
        InputHandler --> FyneApp
    end
    
    style FyneApp fill:#e1f5fe
    style EinkApp fill:#f3e5f5
    style EinkScreen fill:#e8f5e8
    style Folder fill:#fff3e0
```