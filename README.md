# 🧪 Bitácora

> Proyecto desarrollado en conjunto con el **CeFI (Centro de Formación e Investigación)** de la **Facultad de Química – UDELAR (Universidad de la República, Uruguay)**.

---

## 📖 Descripción

**Bitácora** es una aplicación web diseñada para la **gestión de registros de uso de equipamiento** dentro de las instalaciones del CeFI.  
Permite a los usuarios **ingresar nuevos registros** y **visualizarlos** tanto de forma individual como en una **tabla general de registros**.

El sistema fue desarrollado para ser **hospedado localmente**, sin dependencias externas ni recursos distribuidos, garantizando un funcionamiento autónomo dentro de la red del CeFI.

---

## ⚙️ Tecnologías utilizadas

- 🦦 **Go (Golang)** — Backend del servidor web  
- 🧱 **HTML5** — Estructura y vistas  
- 🎨 **CSS3 (inline)** — Estilos embebidos dentro del HTML  
- ⚡ **JavaScript (inline)** — Lógica del frontend embebida  
- 🗃️ **SQLite / almacenamiento local** (dependiendo de la configuración)

> 📝 **Nota importante:**  
> Todo el **JavaScript** y el **CSS** están embebidos directamente en los archivos HTML por requerimiento explícito del cliente, ya que el programa es **hospedado de forma local** en las instalaciones del CeFI.  
> Esto elimina la necesidad de manejar archivos estáticos separados y simplifica la ejecución del sistema en entornos sin conexión.

---

## 🧩 Funcionalidades principales

- ➕ **Alta de registros:** formulario simple para ingresar nuevos usos de equipamiento.  
- 👁️ **Visualización individual:** permite consultar los detalles de un registro específico.  
- 📊 **Visualización en tabla:** listado general de registros ordenados cronológicamente.  
- 🔍 **Búsqueda rápida:** filtro interno por fecha, usuario o equipo (según configuración).  
- 💾 **Ejecución local:** sin necesidad de conexión a Internet ni servidores externos.  

---
