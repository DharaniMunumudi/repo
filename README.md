**# MCP SECURE PLATFORM**

#Project Overview
MCP Secure Platorm consists of 4 services:
1.MCP Server/SSE Transport-Enables real-time communication by streaming updates and responses from the MCP server to connected clients.
2.Gateway Service-Acts as the entry point, routing client requests to the appropriate backend services.
3.Content Serving service-Delivers requested content and resources to users or applications efficiently.
4.Ingestion service-Collects, processes, and ingests data from various sources into the platform for further use.

**#MCP SECURE PLATFORM ARCHITECTURE**

    GitHub / Copilot (MCP Client User)  
          | SSE
     MCP Server(SSE Transport)     
        Owner: Manoj    
          |
     Gateway Service   
      Owner: Dharani    
          |
 -------------------------
 |                        |
AWS Cognito         AWS Secrets Manager   
Authentication      Secrets Storage   
     |
Content Serving Service 
(Owners: Manoj, Swati,Dharani)
     |
PostgreSQL Database 
Vulnerability Content 
Indexed Knowledge Base
     |
Ingestion Service 
 (Owner:Swati)
